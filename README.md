# GophKepper
Менеджер паролей, проект для Яндекс практикум

# Регистрация и аутентификация пользователей

	Используем jwt токены. 
	Регистрация: после сохранения в бд логина и пароля пользователя формируем jwt токен (для signed string используем значение пароля), добавляем в заголовок запроса Authorization.	
	Аутентификация: проверяем значения логина и пароля, если все ок, то формируем jwt токен и добавляем в заголовок запроса Autorization.

