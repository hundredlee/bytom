package contract

import (
	"encoding/json"
	"fmt"
	"github.com/bytom/blockchain/txbuilder"
)

// RevealPreimage stores the information of RevealPreimage contract
type RevealPreimage struct {
	CommonInfo
	Value string `json:"value"`
}

// DecodeRevealPreimage unmarshal JSON-encoded data of contract action
func DecodeRevealPreimage(data []byte) (ContractAction, error) {
	a := new(RevealPreimage)
	err := json.Unmarshal(data, a)
	return a, err
}

// BuildContractReq create new ContractReq which contain contract's name and arguments
func (a *RevealPreimage) BuildContractReq(contractName string) (*ContractReq, error) {
	arguments, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return &ContractReq{
		ContractName: contractName,
		ContractArgs: arguments,
	}, nil
}

// Build create a transaction request
func (a *RevealPreimage) Build() (*string, error) {
	var buildReqStr string

	if a.Alias {
		buildReqStr = fmt.Sprintf(buildAcctRecvReqFmtByAlias, a.OutputID, a.AssetInfo, a.Amount, a.AccountInfo, a.BtmGas, a.AccountInfo)
	} else {
		buildReqStr = fmt.Sprintf(buildAcctRecvReqFmt, a.OutputID, a.AssetInfo, a.Amount, a.AccountInfo, a.BtmGas, a.AccountInfo)
	}

	return &buildReqStr, nil
}

// AddArgs add the parameters for contract
func (a *RevealPreimage) AddArgs(tpl *txbuilder.Template) error {
	if err := addDataArgs(tpl, []string{a.Value}); err != nil {
		return err
	}

	return nil
}