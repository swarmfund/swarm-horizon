package base

//go:generate genny -in=$GOFILE -out=signer_type_flags_populate.go gen "flagValueType=xdr.SignerType"
//go:generate genny -in=$GOFILE -out=asset_policy_type_flags_populate.go gen "flagValueType=xdr.AssetPolicy"
//go:generate genny -in=$GOFILE -out=block_reasons_flags_populate.go gen "flagValueType=xdr.BlockReasons"

func FlagFromflagValueType(mask int32, allFlags []flagValueType) []Flag {
	result := []Flag{}
	for _, flagValue := range allFlags {
		flagValueAsInt := int32(flagValue)
		if (flagValueAsInt & mask) == flagValueAsInt {
			result = append(result, Flag{
				Value: flagValueAsInt,
				Name:  flagValue.String(),
			})
		}
	}

	return result
}
