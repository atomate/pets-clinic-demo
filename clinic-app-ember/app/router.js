import Ember from 'ember';
import config from './config/environment';

var Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('pets', { resetNamespace: true }, function() {
    this.route('add');
    this.route('view', { path: 'view/:pet_id' });
  });
});

export default Router;
