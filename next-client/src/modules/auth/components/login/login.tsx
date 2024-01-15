'use client';

import { apiClient } from '@/modules/api';
import { Button } from '@nextui-org/react';
import { CredentialResponse, GoogleLogin } from '@react-oauth/google';
import { useCallback } from 'react';

export default function Login() {
  const successHandler = useCallback(
    (credentialResponse: CredentialResponse) => {
      apiClient
        .post('/auth/login', {
          token: credentialResponse.credential,
        })
        .then((loginResponse) => {
          console.log(loginResponse.data);
          apiClient.get('/auth/me').then((meResponse) => {
            console.log(meResponse.data);
          });
        });
      console.log(credentialResponse);
    },
    []
  );

  const handleMeCleck = useCallback(() => {
    apiClient.get('/auth/me').then((meResponse) => {
      console.log(meResponse.data);
    });
  }, []);

  return (
    <div>
      <div className={'p-8'}>
        <GoogleLogin
          size={'large'}
          onSuccess={successHandler}
          onError={() => {
            console.log('Login Failed');
          }}
        />
      </div>
      <div className={'p-8'}>
        <Button onClick={handleMeCleck}>Me</Button>
      </div>
    </div>
  );
}
