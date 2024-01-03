
// API呼び出しサンプル
/*
const updateUserPassword = () => {
    const old_user_password = document.getElementById('old_user_password').value;
	const user_password = document.getElementById('user_password').value;
    const user_password_c = document.getElementById('user_password_c').value;

    if (user_password.trim() == '') {
		alert("新パスワードを空白以外の一文字以上を入力して下さい。");
	} else if (user_password !== user_password_c) {
		alert("新パスワードと新パスワード（再）の値が一致していません。");
	} else {
		fetch('/api/account/password', {
            method: 'PUT',
            headers: {"Content-Type": "application/json"},
		    body: JSON.stringify({old_user_password, user_password})
        }).then(response => {
            if (!response.ok) throw new Error(`HTTP Status: ${response.status}`);
            return response.json();

        }).then(data => {
            alert('更新しました。自動でログアウトします。')
            window.location.href = '/logout'

        }).catch(error => {
            if (error.message.includes('Status: 400')) {
                alert("旧パスワードの値が異なります。");
            } else {
                alert("エラーが発生しました。自動でログアウトします。");
                window.location.replace('/logout');
            }
        });
	}
}
*/