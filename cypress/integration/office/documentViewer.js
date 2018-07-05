/* global cy, Cypress */
describe('The document viewer', function() {
  beforeEach(() => {
    Cypress.config('baseUrl', 'http://officelocal:4000');
  });
  it('redirects to sign in when not logged in', function() {
    cy.visit('/moves/foo/documents');
    cy.contains('Welcome');
    cy.contains('Sign In');
  });
  it('produces error when move cannot be found', () => {
    cy.visit('/');
    cy.signInAsUser('9bfa91d2-7a0c-4de0-ae02-b8cf8b4b858b');
    cy.visit('/moves/9bfa91d2-7a0c-4de0-ae02-b8cf8b4b858b/documents');
    cy.contains('An error occurred'); //todo: we want better messages when we are making custom call
  });
  it('loads basic information about the move', () => {
    cy.visit('/');
    cy.signInAsUser('9bfa91d2-7a0c-4de0-ae02-b8cf8b4b858b');
    cy.visit('/moves/F2AF74E2-61B0-40AB-9ABD-172A3863E258/documents');
    cy.contains('Donut, John');
    cy.contains('6TMWRY');
    cy.contains('5789345789');
  });
});

//F2AF74E2-61B0-40AB-9ABD-172A3863E258
