const validate = (form) => {
    const user_name = form.elements['user_name'].value;
    const user_password = form.elements['user_password'].value;
    const user_password_confirm = form.elements['user_password_confirm'].value;

    let error = "";
    if (user_name === "") {
        error = "ユーザ名を入力して下さい。";
    } else if (user_password === "") {
        error = "パスワードを入力して下さい。";
    } else if (user_password !== user_password_confirm) {
		error = "パスワードが一致していません。";
	}

    document.getElementById("error").innerHTML = error;
    return error === "";
}

const signup = () => {
    const form = document.getElementById("signup-form");
    if (!validate(form)) return;

    const user_name = form.elements['user_name'].value;
    const user_password = form.elements['user_password'].value;

    const body = {
        user_name: user_name,
        user_password: user_password
    };

    fetch('/signup', {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(body)
    })
    .then(handleResponse)
    .then(() => {
        window.location.replace('/login');
    })
    .catch(handleError)
    .catch((error) => {
        const status = getErrorStatus(error);
        document.getElementById("error").innerHTML = (status === 409)
        ? "ユーザ名が既に使われています。"
        : "登録に失敗しました。";
    });
}