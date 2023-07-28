""" 

問題: ユーザーのブログ投稿

あなたのタスクは、ユーザーがブログを投稿・管理できるバックエンドアプリケーションを作成することです。
このアプリケーションでは、以下の要件を持つクラスを作成してください：

Userクラス：ユーザーの属性としてユーザー名、メールアドレス、パスワードなどを持ちます。
Postクラス：ブログ投稿の属性としてタイトル、内容、作成日時、ユーザーIDなどを持ちます。
Blogクラス：このクラスは複数の投稿（Postインスタンス）を管理します。投稿の追加、削除、一覧表示などのメソッドを持ちます。
上記の3つのクラスを適切に設計して、以下の点に注意してください：

クラスは適切な属性とメソッドを持ち、適切に相互作用します。
ユーザー名やメールアドレスの重複を防ぐための適切なバリデーションを実装してください。
ユーザーが投稿を追加・削除できるようにメソッドを実装してください。
投稿を一覧表示するメソッドを実装してください。一覧は作成日時の降順で表示されるようにしてください。

"""
import datetime


class User:
    def __init__(self, username, email, password):
        self.username = username
        self.email = email
        self.password = password


class Post:
    def __init__(self, title, content, created_at, user_id):
        self.title = title
        self.content = content
        self.created_at = created_at
        self.user_id = user_id


class Blog:
    def __init__(self):
        self.users = {}
        self.posts = []

    def add_user(self, username, email, password):
        if email in [user.email for user in self.users.values()]:
            raise ValueError("Email already exists.")
        user_id = len(self.users) + 1
        user = User(username, email, password)
        self.users[user_id] = user
        return user_id

    def add_post(self, title, content, user_id):
        if user_id not in self.users:
            raise ValueError("User does not exist.")
        post = Post(title, content, datetime.now(), user_id)
        self.posts.append(post)

    def delete_post(self, post_id):
        for post in self.posts:
            if post.id == post_id:
                self.posts.remove(post)
                break
        else:
            raise ValueError("Post not found.")

    def get_all_posts(self):
        return sorted(self.posts, key=lambda x: x.created_at, reverse=True)
