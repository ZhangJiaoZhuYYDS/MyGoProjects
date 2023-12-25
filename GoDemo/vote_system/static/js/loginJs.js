const registerButton = document.getElementById('register-button');
registerButton.addEventListener('click', () => {
    window.location.href = './register.html';
});

const form = document.querySelector('form');
const result = document.querySelector('.result');

// 获取邮箱，发送验证码请求
const mailInput = document.getElementById('mail');
const sendButton = document.getElementById('send');
sendButton.addEventListener('click', () => {
    const mail = mailInput.value;
    axios.post('http://localhost:8080/sendEmailCode', { mail })
        .then(response => {
            alert(response.data.msg);
        })
        .catch(error => {
            console.error(error);
        });
});



// 点击按钮提交数据
form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries());
    try {
        const response = await axios.post('http://localhost:8080/login', data);
        console.log(response)
        if (response.data.code == 200) {
            window.location.href = "./toupaio.html"
        }
        alert(response.data.msg)
    } catch (error) {
        result.textContent = error.response.data.message;
    }
});