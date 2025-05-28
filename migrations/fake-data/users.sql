-- Очистка таблицы (опционально)
-- TRUNCATE TABLE news.users RESTART IDENTITY;

-- Вставка тестовых пользователей
INSERT INTO news.users (
    email,
    password,
    name,
    created_at
)
VALUES
    ('admin@example.com', 
    '$2a$10$xJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "admin123"
    'Администратор Системы', 
    '2023-01-01 09:00:00'),
    
    ('editor@example.com', 
    '$2a$10$yJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "editor123"
    'Главный Редактор', 
    '2023-01-02 10:15:00'),
    
    ('reporter@example.com', 
    '$2a$10$zJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "reporter123"
    'Журналист Новостей', 
    '2023-01-03 11:30:00'),
    
    ('correspondent@example.com', 
    '$2a$10$aJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "correspondent123"
    'Корреспондент Полевой', 
    '2023-01-04 12:45:00'),
    
    ('photographer@example.com', 
    '$2a$10$bJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "photo123"
    'Фотограф Издательства', 
    '2023-01-05 14:00:00'),
    
    ('analyst@example.com', 
    '$2a$10$cJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "analyst123"
    'Аналитик Данных', 
    '2023-01-06 15:15:00'),
    
    ('intern@example.com', 
    '$2a$10$dJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "intern123"
    'Стажёр Отдела', 
    '2023-01-07 16:30:00'),
    
    ('pr_manager@example.com', 
    '$2a$10$eJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "prmanager123"
    'PR-Менеджер', 
    '2023-01-08 17:45:00'),
    
    ('tech_writer@example.com', 
    '$2a$10$fJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "techwriter123"
    'Технический Писатель', 
    '2023-01-09 19:00:00'),
    
    ('guest@example.com', 
    '$2a$10$gJwL5v5Jz5UZJf5N5YQwUeJ6X9rVd9Xf9Vd9Xf9Vd9Xf9Vd9Xf9Vd', -- хеш пароля "guest123"
    'Гостевой Аккаунт', 
    '2023-01-10 20:15:00');

-- Проверка вставленных данных
SELECT * FROM news.users ORDER BY created_at;