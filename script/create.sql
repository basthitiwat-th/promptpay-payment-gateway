CREATE TABLE IF NOT EXISTS payment_history (
    payment_id INT PRIMARY KEY AUTO_INCREMENT,
    reference_id VARCHAR(255),
    customer_acc_no VARCHAR(25),
    prompt_pay_id VARCHAR(25),
    merchant_id VARCHAR(25),
    amount FLOAT,
    status VARCHAR(20),
    payment_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_reference_id (reference_id)
);