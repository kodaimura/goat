const validateForm = () => {
	const form = document.getElementById("signup-form");
    const user_name = form.elements['user_name'].value;
    const user_password = form.elements['user_password'].value;
    const user_password_confirm = form.elements['user_password_confirm'].value;

    if (user_name === "") {
        document.getElementById("error").innerHTML = "ユーザ名を入力して下さい。"
        return false
    } else if (user_password === "") {
        document.getElementById("error").innerHTML = "パスワードを入力して下さい。"
        return false
    } else if (user_password !== user_password_confirm) {
		document.getElementById("error").innerHTML = "パスワードが一致していません。"
		return false
	} else {
		return true
	}
}

const signup = () => {
    if (!validateForm()) return;

    const form = document.getElementById("signup-form");
    const user_name = form.elements['user_name'].value;
    const user_password = form.elements['user_password'].value;

    const body = {
        user_name: user_name,
        user_password: user_password
    }

    fetch('/signup', {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(body)
    })
    .then(handleResponse)
    .then(() => {
        window.location.replace('/login');
    })
    .catch((error) => {
        const status = error.message.match(/HTTP Status: (\d+)/)[1];
        if (status == "409") {
            document.getElementById("error").innerHTML = "ユーザ名が既に使われています。";
        } else {
            document.getElementById("error").innerHTML = "登録に失敗しました。";
        }
    });
}