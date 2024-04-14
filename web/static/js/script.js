const handleResponse = (response) => {
    if (!response.ok) {
        throw new Error(`HTTP Status: ${response.status}`);
    } else {
        return response.json();
    }
}

const handleError = (error) => {
    const match = error.message.match(/HTTP Status: (\d+)/);
    const status = match? match[1] : "";
    if (status == "500") {
        alert("予期せぬエラーが発生しました。");
        window.location.replace('/logout');
    } else {
        console.error(error);
        return error;
    }
}