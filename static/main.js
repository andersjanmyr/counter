$(function() {
  $('.counter,.logo').click(function() {
    $.post(mountPoint + '/counter', function(data) {
      $('.counter').text(data);
    });
  });
});
