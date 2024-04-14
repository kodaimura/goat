const login = () => {
    const form = document.getElementById("login-form");
    const user_name = form.elements['user_name'].value;
    const user_password = form.elements['user_password'].value;

    const body = {
        user_name: user_name,
        user_password: user_password
    }

    fetch('/login', {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(body)
    })
    .then(handleResponse)
    .then(() => {
        window.location.replace('/');
    })
    .catch((error) => {
        const status = error.message.match(/HTTP Status: (\d+)/)[1];
        if (status == "401") {
            document.getElementById("error").innerHTML = "IDまたはパスワードが異なります。";
        } else {
            document.getElementById("error").innerHTML = "ログインに失敗しました";
        }
    });
}