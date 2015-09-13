$(function() {
  $('.counter,.logo').click(function() {
    $.post('/counter', function(data) {
      $('.counter').text(data);
    });
  });
});
