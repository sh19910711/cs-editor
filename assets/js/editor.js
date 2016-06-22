(function() {
  "use strict";

  window.addEventListener('DOMContentLoaded', _=> {
    document.getElementById('create-project').addEventListener('click', _=> {
      fetch('/api/projects', {method: 'POST'}).then(res => {
        return res.json();
      }).then(json => {
        console.log(json);
      });
    });
  });
})();
