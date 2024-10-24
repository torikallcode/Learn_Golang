package main

type buku struct {
	judul         string
	penulis       string
	jumlahHalaman int
	TahunTerbit   int
}

func TampilkanBuku(buku string) {

}

func main() {

	algoritma := buku{
		judul:         "Algoritma",
		penulis:       "Torikal",
		jumlahHalaman: 300,
		TahunTerbit:   2003,
	}

	golang := buku{
		judul:         "golang",
		penulis:       "Torikal",
		jumlahHalaman: 175,
		TahunTerbit:   2007,
	}

	javaScript := buku{
		judul:         "javaScript",
		penulis:       "Torikal",
		jumlahHalaman: 450,
		TahunTerbit:   2005,
	}

}
