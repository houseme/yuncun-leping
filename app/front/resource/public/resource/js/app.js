// 获取当前页面相关信息
const currentUrl = window.location.href;
const currentHostname = window.location.hostname;
const userAgent = window.navigator.userAgent;
const platform = window.navigator.platform;
const geolocation = new Promise((resolve, reject) => {
    navigator.geolocation.getCurrentPosition(
        (position) => resolve(`${position.coords.latitude},${position.coords.longitude}`),
        (error) => reject(error)
    );
});

const random = Math.random();


// 构建请求参数
const params = new URLSearchParams({
    currentUrl: currentUrl,
    domain: currentHostname,
    platform: platform,
    random: random,
});


const headers = new Headers({
    'Content-Type': 'application/json',
    'Authorization': 'Bearer <token>'
});

const options = {
    method: 'GET',
    headers: headers,
    mode: 'cors',
    // credentials: 'omit',
    cache: 'default'
};
var api_domain = `https://api.comments.haoda.me/api.v1/front/comment?${params}`;
// 发送异步请求
fetch(api_domain, options)
    .then((response) => response.json())
    .then((data) => {
        if (data.code === 200) {
            const commentContent = data.data.comment_content;
            const hotCommentsElement = document.getElementById('hotComments');
            hotCommentsElement.textContent = commentContent;
        } else {
            console.error('Error fetching data:', data.message);
        }
    })
    .catch((error) => {
        console.error('Error fetching data:', error);
    });