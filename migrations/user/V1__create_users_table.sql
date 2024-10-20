-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- Уникальный числовой идентификатор
    keycloak_id UUID UNIQUE NOT NULL, -- UUID из Keycloak
    name VARCHAR(100) NOT NULL, -- Имя пользователя
    surname VARCHAR(100) NOT NULL, -- Фамилия пользователя
    email VARCHAR(100) UNIQUE NOT NULL, -- Email пользователя
    role VARCHAR(50) NOT NULL, -- Роль пользователя
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата обновления
    version INT DEFAULT 1 -- Версия записи для оптимистической блокировки
    );

-- Создаем функцию для автоматического обновления поля updated_at при изменении данных
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создаем триггер, вызывающий функцию перед обновлением записи
CREATE TRIGGER update_users_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

-- Создаем индексы для повышения производительности поиска
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_keycloak_id ON users (keycloak_id);
