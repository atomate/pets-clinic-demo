import Ember from 'ember';

export default Ember.Controller.extend({
    actions: {
        savePet: function(){
            this.get('model').save().then(pet => this.transitionToRoute('pets.view', pet));
        }
    }
});
