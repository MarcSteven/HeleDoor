$(function () {
    var carousel_width = $('.carousel-wrapper img').width();

    setInterval(function () {
        var carousel_left = parseInt($('#ul-carousel').css('margin-left'));
        carousel_left -= carousel_width;
        $('#ul-carousel').animate({'margin-left': carousel_left}, 300);
        if(carousel_left <= -3*carousel_width){
            setTimeout(function () {
                $('#ul-carousel').css('margin-left', 0);
            },400);
        }
    }, 4000);
})

