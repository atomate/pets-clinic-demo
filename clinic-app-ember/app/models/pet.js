import DS from 'ember-data';

export default DS.Model.extend({
  name: DS.attr('string'),
  ownerName: DS.attr('string'),
  ownerPhone: DS.attr('string'),
  visits: DS.hasMany('visit')
});
