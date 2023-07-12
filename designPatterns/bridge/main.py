from abc import ABC, abstractmethod

class MessageSender(ABC):
    def __init__(self, implementation):
        self.implementation = implementation

    @abstractmethod
    def send_message(self, message):
        pass

class MessageSenderImplementation(ABC):
    @abstractmethod
    def send(self, message):
        pass

class SMSMessageSender(MessageSender):
    def send_message(self, message):
        print("Sending SMS message:")
        self.implementation.send(message)

class EmailMessageSender(MessageSender):
    def send_message(self, message):
        print("Sending Email message:")
        self.implementation.send(message)

class DefaultMessageSenderImplementation(MessageSenderImplementation):
    def send(self, message):
        print(f"Default implementation: {message}")

class SecureMessageSenderImplementation(MessageSenderImplementation):
    def send(self, message):
        print(f"Secure implementation: {message}")

def main():
    default_implementation = DefaultMessageSenderImplementation()
    secure_implementation = SecureMessageSenderImplementation()

    sms_sender = SMSMessageSender(default_implementation)
    email_sender = EmailMessageSender(secure_implementation)

    sms_sender.send_message("Hello, this is an SMS!")
    email_sender.send_message("Hello, this is an Email!")

if __name__ == '__main__':
    main()