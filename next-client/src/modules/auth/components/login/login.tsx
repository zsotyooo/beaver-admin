'use client';

import { apiClient } from '@/modules/api';
import { CredentialResponse, GoogleLogin } from '@react-oauth/google';
import { useCallback } from 'react';

export default function Login() {
  const successHandler = useCallback(
    (credentialResponse: CredentialResponse) => {
      apiClient
        .post('/user/login', {
          token: credentialResponse.credential,
        })
        .then((loginResponse) => {
          console.log(loginResponse.data);
          apiClient.get('/user/me').then((meResponse) => {
            console.log(meResponse.data);
          });
        });
      console.log(credentialResponse);
    },
    []
  );

  return (
    <div>
      <GoogleLogin
        size={'large'}
        onSuccess={successHandler}
        onError={() => {
          console.log('Login Failed');
        }}
      />
    </div>
  );
}
