CREATE TABLE history_pemasok (
    id_pemasok INT(11) NOT NULL,
    id_buku INT(11) NOT NULL,
    tanggal DATE NOT NULL,
    stok INT(11) NOT NULL,
    ket ENUM('Buku Baru', 'Tambah Stok'),
    FOREIGN KEY (id_pemasok) REFERENCES pemasok(id), 
    FOREIGN KEY (id_buku) REFERENCES buku(id)  
) ENGINE = InnoDB;