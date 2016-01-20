import DS from 'ember-data';

export default DS.Model.extend({
  pet: DS.belongsTo("pet"),
  result: DS.attr('string')
});
