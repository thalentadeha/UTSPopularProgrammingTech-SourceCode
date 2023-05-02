package datadiri

type DataDiriOrang struct {
	Nama   string
	Nim    int
	Alamat string
}

func ShowDataDiri(nama string) DataDiriOrang {

	DataOrang := DataDiriOrang{
		Nama:   "not found",
		Nim:    0,
		Alamat: "not found",
	}

	if nama == "gusti" {
		DataOrang := DataDiriOrang{
			Nama:   "Gusti",
			Nim:    1234,
			Alamat: "Malang",
		}

		return DataOrang
	} else if nama == "pangestu" {
		DataOrang := DataDiriOrang{
			Nama:   "Pangestu",
			Nim:    321,
			Alamat: "Kabupaten Malang",
		}

		return DataOrang
	}

	return DataOrang

}
