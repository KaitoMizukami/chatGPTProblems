"""

問題: 銀行の口座管理システムを開発しています。システム内で唯一の銀行データベース接続オブジェクトを使用する必要があります。
Singletonデザインパターンを使用して、銀行データベース接続オブジェクトの唯一のインスタンスを保証するクラスを実装してください。

以下の要件に基づいて、Singletonデザインパターンを使用して問題を解決してください。

BankDatabaseConnectionクラスは、システム内で唯一のインスタンスを持つ必要があります。同じデータベース接続を複数の場所で使用することはできません。

"""


class BankDatabaseConnection:
    _instance = None

    def __new__(cls):
        if not cls._instance:
            cls._instance = super(BankDatabaseConnection, cls).__new__(cls)
        return cls._instance

    @staticmethod
    def get_instance(cls):
        return cls._instance
            


def main():
    # BankDatabaseConnectionのインスタンスを取得
    connection1 = BankDatabaseConnection()
    connection2 = BankDatabaseConnection()

    # インスタンスが同じか確認
    print(connection1 is connection2)  # True

if __name__ == '__main__':
    main()