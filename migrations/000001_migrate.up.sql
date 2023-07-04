
CREATE TABLE goods (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(45),
                       sort VARCHAR(45)
);

CREATE TABLE client (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        name VARCHAR(45),
                        created DATE
);

CREATE TABLE purchase (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          name VARCHAR(45),
                          client_id INT,
                          FOREIGN KEY (client_id) REFERENCES client(id)
);

CREATE TABLE purchase_goods (
                                id INT AUTO_INCREMENT PRIMARY KEY,
                                purchase_id INT,
                                FOREIGN KEY (purchase_id) REFERENCES purchase(id),
                                goods_id INT,
                                FOREIGN KEY (goods_id) REFERENCES goods(id),
                                amount DECIMAL,
                                cort_price INT
);

CREATE TABLE requirement (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             date DATE,
                             client_id INT,
                             FOREIGN KEY (client_id) REFERENCES client(id)
);

CREATE TABLE requirement_goods (
                                   id INT AUTO_INCREMENT PRIMARY KEY,
                                   requirement_id INT,
                                   FOREIGN KEY (requirement_id) REFERENCES requirement(id),
                                   goods_id INT,
                                   FOREIGN KEY (goods_id) REFERENCES goods(id),
                                   amount DECIMAL,
                                   cost_cell INT
);
