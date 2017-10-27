package name_service;

import "context"
import "math/big"
import "os"
import "github.com/ethereum/go-ethereum/accounts/abi/bind"
import "github.com/ethereum/go-ethereum/common"
import "github.com/ethereum/go-ethereum/crypto"

type server struct {
	contract *NameService
}

func NewServer(address common.Address, backend bind.ContractBackend) NameServiceServer {
	contract, _ := NewNameService(address, backend)
	return &server{
		contract: contract,
	}
}


func (s *server) GetName(c context.Context, r *EmptyReq) (*GetNameResp, error) {
	data, err := s.contract.GetName(
		&bind.CallOpts{
			Pending: true,
			Context: c,
		},
	)
	return &GetNameResp{ 
		Arg: data,
	}, err
}


func (s *server) SetName(c context.Context, r *SetNameReq) (*TransactionResp, error) {
	tx, err := s.contract.SetName(
		r.GetOpts().TransactOpts(),
		r.Name,
	)
	if tx == nil {
		return nil, err
	}
	return &TransactionResp{
		Hash: tx.Hash().Hex(),
	}, err
}


// TransactOpts converts to bind.TransactOpts
func (m *TransactOpts) TransactOpts() *bind.TransactOpts {
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(m.PrivateKey))
	if err != nil {
		os.Exit(-1)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = big.NewInt(4712388)
	auth.GasPrice = big.NewInt(20000000000)
	auth.Nonce = nil //big.NewInt(m.Nonce)
	auth.Value = big.NewInt(m.Value)
	return auth
}
