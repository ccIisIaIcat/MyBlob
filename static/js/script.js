const sidebar = document.getElementById('sidebar');

// 当鼠标悬停在导航栏上时展开
sidebar.addEventListener('mouseenter', function() {
    this.style.left = '0';
});

// 当鼠标离开导航栏时隐藏
sidebar.addEventListener('mouseleave', function() {
    this.style.left = '-250px';
});