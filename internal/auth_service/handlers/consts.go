package handlers

import "time"

const (
	REQUEST_TIMEOUT                    = 10 * time.Second
	TOKEN_EXPIRATION_TIME              = 2 * 24 * time.Hour
	NOTIFICATION_MESSAGE_SUBJECT       = "IP-адрес изменен!"
	NOTIFICATION_MESSAGE_BODY_TEMPLATE = ("Уважаемый %s!\r\n" +
		"В ваш аккаунт был произведен вход с другого IP-адреса (%s).\r\n" +
		"Если это были не вы - немедленно поменяйте данные для авторизации!\r\n\r\n" +
		"С уважением, Demo Auth Service.")
)
