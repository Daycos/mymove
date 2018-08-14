/* global cy */

// Sign in user
//  choose hhg move type
// Fill in the form (including optional addresses)
// click next (which submits data)
// verify next page is Review page
// Update hhg. . .
describe('completing the hhg flow', function() {
  beforeEach(() => {
    // sm_hhg@example.com
    cy.signInAsUser('4b389406-9258-4695-a091-0bf97b5a132f');
  });

  it('selects hhg and progresses thru form', function() {
    cy.visit('localhost:4000/moves/8718c8ac-e0c6-423b-bdc6-af971ee05b9a');

    cy.contains('Household Goods Move').click();
    cy.contains('Next').click({ force: true });
    cy.location().should(loc => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/hhg-start/);
    });
    // cy.get(calendar date)
    // Pickup address
    cy
      .get('input[name="pickup_address.street_address_1"]')
      .clear({ force: true })
      .type('123 Elm Street');
    cy
      .get('input[name="pickup_address.city"]')
      .clear()
      .type('Alameda');
    cy.get('select[name="pickup_address.state"]').select('CA');
    cy.get('input[name="pickup_address.postal_code"]').type('94607');
    cy.get('input[type="radio"]').check('yes', { force: true }); // checks yes for both radios on form

    // Secondary pickup address
    cy
      .get('input[name="secondary_pickup_address.street_address_1"]')
      .clear({ force: true })
      .type('543 Oak Street');
    cy
      .get('input[name="secondary_pickup_address.city"]')
      .clear()
      .type('Oakland');
    cy.get('select[name="secondary_pickup_address.state"]').select('CA');
    cy.get('input[name="secondary_pickup_address.postal_code"]').type('94609');

    // Destination address
    cy
      .get('input[name="delivery_address.street_address_1"]')
      .clear({ force: true })
      .type('678 Madrone Street');
    cy
      .get('input[name="delivery_address.city"]')
      .clear()
      .type('Fremont');
    cy.get('select[name="delivery_address.state"]').select('CA');
    cy.get('input[name="delivery_address.postal_code"]').type('94567');

    // Weights
    cy
      .get('input[name="weight_estimate"]')
      .type('3000')
      .get('input[name="progear_weight_estimate"]')
      .type('250')
      .get('input[name="spouse_progear_weight_estimate"]')
      .type('150');
  });
});
