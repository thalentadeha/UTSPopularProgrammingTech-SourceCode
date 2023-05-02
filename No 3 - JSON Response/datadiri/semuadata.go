package datadiri

type SemuaDataDiri struct {
	Data1 DataDiriOrang
	Data2 DataDiriOrang
}

func ShowSemuaData() SemuaDataDiri {

	semuadata := SemuaDataDiri{
		Data1: ShowDataDiri("gusti"),
		Data2: ShowDataDiri("pangestu"),
	}

	return semuadata

}
