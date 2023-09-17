# GophKeeper
Менеджер паролей, проект для Яндекс практикум

### Регистрация и аутентификация пользователей

Используем jwt токены. 
- Регистрация: после сохранения в бд логина и пароля пользователя формируем jwt токен (для signed string используем значение пароля), добавляем в заголовок запроса Authorization.
- Аутентификация: проверяем значения логина и пароля, если все ок, то формируем jwt токен и добавляем в заголовок запроса Autorization.

### Хранение данных

БД PostgreSQL

![image](https://github.com/kartalenka7/GophKeeper/assets/113780951/b54c1cae-d164-445c-bb84-389c2a5db9f6)

#### Безопасность

- Пароль пользователя переводится в хэш и записывается в бд в виде хэша
sha256.Sum256([]byte(password)), при проверке пароля вычисляем хэш переданного пароля и сравниваем с хэшем из бд.

- Также шифруем поле data перед записью в таблицу DataTable с использованием алгоритмов GCM и AES256:
  1) в качестве ключа шифрования key используем sha256 хэш пароля, получаем cipher.Block
   ###### block, err := aes.NewCipher(key)
  2) Создаем GCM режим шифрования
   ###### aesGCM, err := cipher.NewGCM(block)
  3) создаём вектор инициализации из последних aesgcm.NonceSize() байт ключа
   ###### iv := key[len(key)-aesGCM.NonceSize():]
  4) зашифровываем, в этом виде будем сохранять в бд
   ###### dst := aesGCM.Seal(nil, iv, data, nil) 

### Протокол взаимодействия клиента и сервера

протокол gRPC

![Диаграмма классов (9)](https://github.com/kartalenka7/GophKeeper/assets/113780951/d4a29544-b45d-4a35-870b-1a62ce3b1c9a)



