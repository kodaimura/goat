const login = () => {
    const form = document.getElementById("login-form");
    const user_name = form.elements['user_name'].value;
    const user_password = form.elements['user_password'].value;

    const body = {
        user_name: user_name,
        user_password: user_password
    };

    fetch('/api/login', {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(body)
    })
    .then(handleResponse)
    .then(() => {
        window.location.replace('/');
    })
    .catch(handleError)
    .catch((error) => {
        const status = getErrorStatus(error);
        document.getElementById("error").innerHTML = (status === 401)
        ? "ユーザ名またはパスワードが異なります。" 
        : "ログインに失敗しました。";
    });
}