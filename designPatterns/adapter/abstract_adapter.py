from abc import ABCMeta, abstractmethod


class EmailAdapter(metaclass=ABCMeta):
    @abstractmethod
    def send_email(self, to, subject, body):
        pass


class SendEmail(EmailAdapter):
    def __init__(self, sender):
        self.__sender = sender 

    def send_email(self, to, subject, body):
        self.__sender.send(to, subject, body)


class EmailSender:
    def send_email(self, to, subject, body):
        print('EmailSenderでメールを送る')
        print(f'{to}, {subject}, {body}')


class NewEmaiLib(EmailAdapter):
    def send(self, recipient, subject, message):
        print('NewEmailLibでメールを送る')
        print(f'{recipient}, {subject}, {message}')
