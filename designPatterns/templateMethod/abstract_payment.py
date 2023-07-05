from abc import ABCMeta, abstractmethod


class AbstractPayment(metaclass=ABCMeta):
    @abstractmethod
    def input_payment_info(self):
        pass 

    @abstractmethod
    def validate_payment_info(self):
        pass 

    @abstractmethod
    def execute_payment(self):
        pass 

    @abstractmethod
    def display(self):
        pass 

    def process_payment(self):
        self.input_payment_info()
        self.validate_payment_info()
        self.execute_payment()
        self.display()


class CreditCardPayment(AbstractPayment):
    def input_payment_info(self):
        print('カード番号、有効期限、セキュリティコードの入力')

    def validate_payment_info(self):
        print('カード番号、有効期限、セキュリティコードの検証')

    def execute_payment(self):
        print('クレジットカードの支払い処理')

    def display(self):
        print('支払い成功or失敗')


class PayPalPayment(AbstractPayment):
    def input_payment_info(self):
        print('PayPalアカウントのメールアドレスとパスワードの入力')

    def validate_payment_info(self):
        print('PayPalアカウントのメールアドレスとパスワードの検証')

    def execute_payment(self):
        print('PayPalの支払い処理')

    def display(self):
        print('支払い成功or失敗')