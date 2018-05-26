import {
  UPDATE_SERVICE_MEMBER,
  GET_SERVICE_MEMBER,
  CREATE_SERVICE_MEMBER,
  serviceMemberReducer,
} from './ducks';

import { GET_LOGGED_IN_USER } from 'shared/User/ducks';
import { get } from 'lodash';
import loggedInUserPayload from 'shared/user/sampleLoggedInUserPayload';
const smPayload = { ...loggedInUserPayload.payload.service_member };
const expectedSM = {
  affiliation: 'ARMY',
  backup_contacts: [],
  backup_mailing_address: {
    city: 'Washington',
    postal_code: '20021',
    state: 'DC',
    street_address_1: '200 K St',
  },
  created_at: '2018-05-25T15:48:49.918Z',
  current_station: {
    address: {
      city: 'Colorado Springs',
      country: 'United States',
      postal_code: '80913',
      state: 'CO',
      street_address_1: 'n/a',
    },
    affiliation: 'ARMY',
    created_at: '2018-05-20T18:36:45.034Z',
    id: '28f63a9d-8fff-4a0f-84ef-661c5c8c354e',
    name: 'Ft Carson',
    updated_at: '2018-05-20T18:36:45.034Z',
  },
  edipi: '1234567890',
  email_is_preferred: false,
  first_name: 'Erin',
  has_social_security_number: true,
  id: '1694e00e-17ff-43fe-af6d-ab0519a18ff2',
  is_profile_complete: true,
  last_name: 'Stanfill',
  middle_name: '',
  personal_email: 'erin@truss.works',
  phone_is_preferred: true,
  rank: 'O_4_W_4',
  residential_address: {
    city: 'Somewhere',
    postal_code: '80913',
    state: 'CO',
    street_address_1: '123 Main',
  },
  telephone: '555-555-5556',
  text_message_is_preferred: true,
  updated_at: '2018-05-25T21:39:10.484Z',
  user_id: 'b46e651e-9d1c-4be5-bb88-bba58e817696',
};
describe('Service Member Reducer', () => {
  describe('GET_LOGGED_IN_USER', () => {
    it('should handle SUCCESS', () => {
      const newState = serviceMemberReducer({}, loggedInUserPayload);
      expect(newState).toEqual({
        currentServiceMember: { ...expectedSM },
        hasLoadError: false,
        hasLoadSuccess: true,
      });
    });
  });
  describe('UPDATE_SERVICE_MEMBER', () => {
    it('Should handle UPDATE_SERVICE_MEMBER_SUCCESS', () => {
      const initialState = { currentServiceMember: null };
      const newState = serviceMemberReducer(initialState, {
        type: UPDATE_SERVICE_MEMBER.success,
        payload: smPayload,
      });

      expect(newState).toEqual({
        currentServiceMember: expectedSM,
        hasSubmitError: false,
        hasSubmitSuccess: true,
      });
    });

    it('Should handle UPDATE_SERVICE_MEMBER_FAILURE', () => {
      const initialState = { currentServiceMember: { id: 'bad' } };

      const newState = serviceMemberReducer(initialState, {
        type: UPDATE_SERVICE_MEMBER.failure,
        error: 'No bueno.',
      });

      expect(newState).toEqual({
        currentServiceMember: { id: 'bad' },
        hasSubmitError: true,
        hasSubmitSuccess: false,
        error: 'No bueno.',
      });
    });
  });
  describe('GET_SERVICE_MEMBER', () => {
    it('Should handle GET_SERVICE_MEMBER_SUCCESS', () => {
      const initialState = { currentServiceMember: null };
      const newState = serviceMemberReducer(initialState, {
        type: GET_SERVICE_MEMBER.success,
        payload: smPayload,
      });

      expect(newState).toEqual({
        currentServiceMember: expectedSM,
        hasSubmitError: false,
        hasSubmitSuccess: true,
      });
    });

    it('Should handle GET_SERVICE_MEMBER_FAILURE', () => {
      const initialState = { currentServiceMember: null };

      const newState = serviceMemberReducer(initialState, {
        type: GET_SERVICE_MEMBER.failure,
        error: 'No bueno.',
      });

      expect(newState).toEqual({
        currentServiceMember: null,
        hasSubmitError: true,
        hasSubmitSuccess: false,
        error: 'No bueno.',
      });
    });
  });
  describe('CREATE_SERVICE_MEMBER', () => {
    it('Should handle CREATE_SERVICE_MEMBER_SUCCESS', () => {
      const initialState = { currentServiceMember: null };
      const newState = serviceMemberReducer(initialState, {
        type: CREATE_SERVICE_MEMBER.success,
        payload: smPayload,
      });

      expect(newState).toEqual({
        currentServiceMember: expectedSM,
        hasSubmitError: false,
        hasSubmitSuccess: true,
        isLoading: false,
      });
    });

    it('Should handle CREATE_SERVICE_MEMBER_FAILURE', () => {
      const initialState = { currentServiceMember: null };

      const newState = serviceMemberReducer(initialState, {
        type: CREATE_SERVICE_MEMBER.failure,
        error: 'No bueno.',
      });

      expect(newState).toEqual({
        currentServiceMember: null,
        hasSubmitError: true,
        hasSubmitSuccess: false,
        error: 'No bueno.',
        isLoading: false,
      });
    });
  });
});
