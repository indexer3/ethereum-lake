package constant

import "github.com/indexer3/ethereum-lake/common/utils"

var (
	erc721MetadataInterfaceID = make([]byte, 4)
	erc721InterfaceID         = make([]byte, 4)

	erc1155InterfaceID = make([]byte, 4)
)

func init() {
	erc721MetadataInterfaceID = utils.XorResultMethods(erc721MethodSigMust)
	erc721InterfaceID = utils.XorResultMethods(erc721MethodSigMust)

	erc1155InterfaceID = utils.XorResultMethods(erc1155SigMust)
}

func ERC721MetadataInterfaceID() [4]byte {
	temp := make([]byte, len(erc721MetadataInterfaceID))
	copy(temp, erc721MetadataInterfaceID)
	return [4]byte{temp[0], temp[1], temp[2], temp[3]}
}

func ERC721InterfaceID() [4]byte {
	temp := make([]byte, len(erc721InterfaceID))
	copy(temp, erc721InterfaceID)
	return [4]byte{temp[0], temp[1], temp[2], temp[3]}
}

func ERC1155InterfaceID() [4]byte {
	temp := make([]byte, len(erc1155InterfaceID))
	copy(temp, erc1155InterfaceID)
	return [4]byte{temp[0], temp[1], temp[2], temp[3]}
}
