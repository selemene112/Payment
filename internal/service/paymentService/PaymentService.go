package paymentservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"payment/internal/entity"
	"payment/internal/excaption"
	"payment/internal/helper"
	paymentmodel "payment/internal/model/paymentModel"
	"payment/internal/repository"
	"payment/pkg/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PaymentTransaction(request paymentmodel.PaymentRequest, c *gin.Context) (paymentmodel.PaymentResponse, error) {
	fmt.Println(request, "Request Gua ini Bang")
	err := helper.ValidateInput(request, c)
	if err != nil {
		return paymentmodel.PaymentResponse{}, err
	}
	token, err := helper.GetToken(c)
	if err != nil {
		return paymentmodel.PaymentResponse{}, err
	}
	wallet := entity.Wallet{}
	walletUpdate := entity.Wallet{}
	walletTransfer := entity.Wallet{}
	walletTransferUpdate := entity.Wallet{}
	var transactionStatus entity.TransactionStatus = entity.TransactionStatusSuccess

	err = config.GetPgsqlDB().Transaction(func(tx *gorm.DB) error {
		if err := repository.GetWalletByUserId(&wallet, token.Id, tx, c); err != nil {
			return err
		}

		// Seleksi data Transaksi
		switch request.TransactionType {
		//transaksi topup
		case entity.TransactionTypeTopup:
			fmt.Println("di Topup")

			amountUpdate := wallet.Amount + request.Amount
			if err := repository.UpdateWalletAmount(&walletUpdate, wallet.ID, amountUpdate, tx, c); err != nil {
				transactionStatus = entity.TransactionStatusFailed
				return err
			}
			if walletUpdate.Amount-wallet.Amount != request.Amount {
				excaption.ErrorHandler(c, http.StatusConflict, errors.New("Race Condition detected"), "Race Condition detected")
				transactionStatus = entity.TransactionStatusFailed
				return fmt.Errorf("race condition detected: %w", err)
			}
			fmt.Println(walletUpdate, "Wallet Update")
			//transaksi withdraw
		case entity.TransactionTypeWithdraw:
			fmt.Println("di Withdraw")
			fmt.Println(request.Amount, wallet.Amount, "Amount Withdraw")
			if request.Amount > wallet.Amount {
				excaption.ErrorHandler(c, http.StatusBadRequest, errors.New("Not enough money"), "Not enough money")
				transactionStatus = entity.TransactionStatusFailed
				return errors.New("Not enough money")
			}
			amountUpdate := wallet.Amount - request.Amount
			if err := repository.UpdateWalletAmount(&walletUpdate, wallet.ID, amountUpdate, tx, c); err != nil {
				transactionStatus = entity.TransactionStatusFailed
				return err
			}
			if walletUpdate.Amount+request.Amount != wallet.Amount {
				excaption.ErrorHandler(c, http.StatusConflict, errors.New("Race Condition detected"), "Race Condition detected")
				transactionStatus = entity.TransactionStatusFailed
				return fmt.Errorf("race condition detected: %w", err)
			}

			//transaksi transfer
		case entity.TransactionTypeTransfer:
			fmt.Println("di Transfer")
			if request.Amount > wallet.Amount {
				excaption.ErrorHandler(c, http.StatusBadRequest, errors.New("Not enough money"), "Not enough money")
				transactionStatus = entity.TransactionStatusFailed
				return fmt.Errorf("not enough money: %w", err)
			}
			if request.DestinationWalletId == (uuid.UUID{}) {
				excaption.ErrorHandler(c, http.StatusBadRequest, errors.New("Invalid Destination Wallet Id"), "Invalid Destination Wallet Id")
				transactionStatus = entity.TransactionStatusFailed
				return fmt.Errorf("invalid destination wallet id: %w", err)
			}

			amountUpdate := wallet.Amount - request.Amount
			if err := repository.UpdateWalletAmount(&walletUpdate, wallet.ID, amountUpdate, tx, c); err != nil {
				transactionStatus = entity.TransactionStatusFailed
				return err
			}
			if walletUpdate.Amount+request.Amount != wallet.Amount {
				excaption.ErrorHandler(c, http.StatusConflict, errors.New("Race Condition detected"), "Race Condition detected")
				transactionStatus = entity.TransactionStatusFailed
				return fmt.Errorf("race condition detected: %w", err)
			}

			if err = repository.GetWalletById(&walletTransfer, request.DestinationWalletId, tx, c); err != nil {
				transactionStatus = entity.TransactionStatusFailed
				return err
			}
			amountTransferUpdate := walletTransfer.Amount + request.Amount
			if err = repository.UpdateWalletAmount(&walletTransferUpdate, walletTransfer.ID, amountTransferUpdate, tx, c); err != nil {
				transactionStatus = entity.TransactionStatusFailed
				return err
			}
		case entity.TransactionTypeCheckBalance:
			if err := repository.GetWalletByUserId(&wallet, token.Id, tx, c); err != nil {
				transactionStatus = entity.TransactionStatusFailed
				return err
			}

		default:
			excaption.ErrorHandler(c, http.StatusBadRequest, errors.New("Invalid Transaction Type"), "Invalid Transaction Type")
			transactionStatus = entity.TransactionStatusFailed
			return fmt.Errorf("invalid transaction type: %w", err)
		}
		return nil
	})
	if err != nil {
		transactionStatus = entity.TransactionStatusFailed
	}
	response := paymentmodel.PaymentResponse{
		UserId:              token.Id,
		WalletId:            wallet.ID,
		TransactionDate:     time.Now(),
		BalanceBefore:       wallet.Amount,
		Amount:              request.Amount,
		BalanceAfter:        walletUpdate.Amount,
		TransactionType:     request.TransactionType,
		DestinationWalletId: request.DestinationWalletId,
	}

	requestJSON, _ := json.Marshal(request)
	responseJSON, _ := json.Marshal(response)

	transaction := entity.Transaction{
		UserId:              token.Id,
		WalletId:            wallet.ID,
		DestinationWalletId: request.DestinationWalletId,
		Request:             json.RawMessage(requestJSON),
		Response:            json.RawMessage(responseJSON),
		Amount:              request.Amount,
		TransactionType:     request.TransactionType,
		TransactionStatus:   transactionStatus,
	}
	recordErr := repository.CreateTransaction(&transaction, c)
	if recordErr != nil {
		return paymentmodel.PaymentResponse{}, recordErr
	}
	if err != nil {
		return paymentmodel.PaymentResponse{}, err
	}

	return response, nil
}
