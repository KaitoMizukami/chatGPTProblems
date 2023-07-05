"""
問題: オンラインショッピングアプリケーションでは、商品を購入するための支払い方法が複数あります。
支払い方法にはクレジットカード、PayPalなどがあります。
それぞれの支払い方法は異なる手順を必要としますが、支払いプロセスは共通のフローを持っています。

以下の要件に基づいて、支払いプロセスを実装するためにTemplate Methodデザインパターンを使用してください。

すべての支払い方法には、次の手順を含む共通のフローがあります:
a. 支払い情報の入力
b. 支払い情報の検証
c. 支払いの実行
d. 支払い結果の表示

クレジットカードの支払い方法では、支払い情報の入力にはカード番号、有効期限、セキュリティコードが必要です。
PayPalの支払い方法では、支払い情報の入力にはPayPalアカウントのメールアドレスとパスワードが必要です。

上記の要件に基づいて、Template Methodデザインパターンを使用して、支払いプロセスを実装してください。

この例では、共通の支払いフローを抽象クラスとして実装し、具体的な支払い方法はサブクラスとして実装します。
各支払い方法の具体的な手順は抽象メソッドとして定義し、サブクラスでオーバーライドします。
そして、抽象クラスのテンプレートメソッドを呼び出すことで、支払いプロセス全体が実行されるようにします。
"""

from abstract_payment import CreditCardPayment, PayPalPayment


def main():
    # クレジットカード支払い
    credit_card_payment = CreditCardPayment()
    credit_card_payment.process_payment()

    # PayPal支払い
    paypal_payment = PayPalPayment()
    paypal_payment.process_payment()

if __name__ == '__main__':
    main()