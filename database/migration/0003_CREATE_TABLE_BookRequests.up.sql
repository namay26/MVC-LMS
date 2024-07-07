CREATE TABLE BookRequests (
    RequestID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT,
    BookID INT,
    RequestDate TIMESTAMP NOT NULL DEFAULT NOW(),
    AcceptDate TIMESTAMP DEFAULT NULL,
    ReturnDate TIMESTAMP DEFAULT NULL,
    Status ENUM('Pending','Approved','Returned') DEFAULT 'Returned' NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(userid),
    FOREIGN KEY (BookID) REFERENCES books(id)
);