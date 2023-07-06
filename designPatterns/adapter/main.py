"""

問題: 外部のメール送信ライブラリを使用して、メールを送信する機能を開発しています。
しかし、既存のコードベースでは別のメール送信インターフェースを使用しており、新しいライブラリとの互換性がありません。
Adapterデザインパターンを使用して、既存のコードベースと新しいライブラリを統合するためのクラスを作成してください。

以下の要件に基づいて、Adapterデザインパターンを使用して問題を解決してください。

既存のコードベースでは、EmailSenderという名前のクラスがメールの送信を処理します。このクラスはsend_emailメソッドを持ち、to, subject, bodyなどのパラメータを受け取ります。
新しいメール送信ライブラリはNewEmailLibというクラスで提供されます。このクラスはsendメソッドを持ち、recipient, subject, messageなどのパラメータを受け取ります。
既存のコードベースのEmailSenderクラスを変更せずに、NewEmailLibクラスを使用してメールを送信するためのアダプタクラスを作成してください。アダプタクラスはEmailSenderクラスとNewEmailLibクラスの間でメールの送信を変換する役割を果たします。

上記の要件に基づいて、Adapterデザインパターンを使用して問題を解決してください。Pythonのクラスを使用してアダプタクラスを作成し、既存のコードベースと新しいメール送信ライブラリを統合してください。

"""
from abstract_adapter import SendEmail, NewEmaiLib


def main():
    adapter = SendEmail(NewEmaiLib)

    adapter.send_email('You', 'About the graduation', 'This is called by NewEmailLib')

if __name__ == '__main__':
    main()