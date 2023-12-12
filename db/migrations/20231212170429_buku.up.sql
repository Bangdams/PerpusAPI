CREATE TABLE buku (
    id INT NOT NULL AUTO_INCREMENT,
    nama VARCHAR(150) NOT NULL,
    penerbit_id INT(11) NOT NULL,
    kategori INT(11) NOT NULL,
    stok INT(11) NOT NULL,
    PRIMARY KEY (id), 
    FOREIGN KEY (penerbit_id) REFERENCES penerbit(id), 
    FOREIGN KEY (kategori) REFERENCES kategori(id)  
) ENGINE = InnoDB;