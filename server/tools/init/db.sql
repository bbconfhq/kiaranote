CREATE DATABASE kiaranote DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
CREATE DATABASE kiaranote_test DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
CREATE USER 'kiara_admin'@'%' IDENTIFIED BY 'kiara_admin_pass';
GRANT ALL PRIVILEGES ON kiaranote.* TO 'kiara_admin'@'%';
GRANT ALL PRIVILEGES ON kiaranote_test.* TO 'kiara_admin'@'%';