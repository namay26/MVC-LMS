CREATE TABLE Users (
    userid INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    pass VARCHAR(255) NOT NULL,
    isAdmin BOOLEAN DEFAULT 0 NOT NULL,
    acctcreate TIMESTAMP DEFAULT NOW() NOT NULL,
    adminStatus ENUM('NotRequested','Pending','isAdmin') DEFAULT 'NotRequested' NOT NULL
);