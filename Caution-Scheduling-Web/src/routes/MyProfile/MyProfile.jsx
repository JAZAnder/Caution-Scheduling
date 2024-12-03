import React, { useState, useEffect } from 'react';
import Background from '../../background';
import './MyProfile.css';

function MyProfile() {
  const [userData, setUserData] = useState(null);
  const [changePasswordVisible, setChangePasswordVisible] = useState(false);
  const [oldPassword, setOldPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [passwordChangeMessage, setPasswordChangeMessage] = useState('');
  const [isSuccessMessage, setIsSuccessMessage] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  const baseUrl = '/api';

  useEffect(() => {
    async function fetchUserData() {
      try {
        const response = await fetch(`${baseUrl}/luser/whoami`, {
          method: 'GET',
          credentials: 'include',
        });
        if (response.ok) {
          const data = await response.json();
          setUserData(data);
        } else {
          setError(true);
        }
      } catch (err) {
        console.error('Error fetching user data:', err);
        setError(true);
      } finally {
        setLoading(false);
      }
    }
    fetchUserData();
  }, []);

  const handleChangePassword = async (e) => {
    e.preventDefault();
    setPasswordChangeMessage('');
    setIsSuccessMessage(false);

    const capitalLetterRegex = /[A-Z]/;
    const specialCharRegex = /[!@#$%^&*(),.?":{}|<>]/;

    if (!capitalLetterRegex.test(newPassword) || !specialCharRegex.test(newPassword)) {
      setPasswordChangeMessage(
        'New password must contain at least one capital letter and one special character.'
      );
      setIsSuccessMessage(false);
      return;
    }

    const params = new URLSearchParams();
    params.append('oldPassword', oldPassword);
    params.append('newPassword', newPassword);

    try {
      const response = await fetch(`${baseUrl}/luser/resetmypasswd`, {
        method: 'PUT',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: params.toString(),
      });

      if (response.ok) {
        setPasswordChangeMessage('Password changed successfully.');
        setIsSuccessMessage(true);
        setOldPassword('');
        setNewPassword('');
      } else {
        const errorData = await response.json();
        if (errorData.error === 'Password is Incorrect') {
          setPasswordChangeMessage('Old password is incorrect.');
        } else {
          setPasswordChangeMessage(`Failed to change password: ${errorData.error || 'Unknown error'}`);
        }
        setIsSuccessMessage(false);
      }
    } catch (err) {
      console.error('Error changing password:', err);
      setPasswordChangeMessage('Failed to change password: Network error.');
      setIsSuccessMessage(false);
    }
  };

  if (loading) {
    return (
      <center>
        <div className="loader"></div>
      </center>
    );
  }

  if (error || !userData) {
    return <div className="text-center">Failed to load user data.</div>;
  }

  return (
    <>
      <Background />
      <div className="my-profile-container">
        <div className="myprofile-body">
          <div className="myprofile-page">
            <h2 className="myprofile-h2">Personal Information</h2>
            <div className="myprofile-info">
              <div className="myprofile-info-column">
                <p><strong>First Name</strong></p>
                <p>{userData.firstName || 'N/A'}</p>
              </div>
              <div className="myprofile-info-column">
                <p><strong>Last Name</strong></p>
                <p>{userData.lastName || 'N/A'}</p>
              </div>
            </div>
            <h2 className="myprofile-h2">Public Information</h2>
            <div className="myprofile-info">
              <div className="myprofile-info-column">
                <p><strong>Email</strong></p>
                <p>{userData.email || 'N/A'}</p>
              </div>
              <div className="myprofile-info-column">
                <p><strong>Role</strong></p>
                <p>{userData.role || 'N/A'}</p>
              </div>
            </div>
            <button
              className="change-password-button"
              onClick={() => setChangePasswordVisible(!changePasswordVisible)}
            >
              {changePasswordVisible ? 'Cancel' : 'Change Password'}
            </button>
            {changePasswordVisible && (
              <form className="change-password-form" onSubmit={handleChangePassword}>
                <div className="form-group">
                  <label>Old Password:</label>
                  <input
                    type="password"
                    value={oldPassword}
                    onChange={(e) => setOldPassword(e.target.value)}
                    required
                  />
                </div>
                <div className="form-group">
                  <label>New Password:</label>
                  <input
                    type="password"
                    value={newPassword}
                    onChange={(e) => setNewPassword(e.target.value)}
                    required
                  />
                </div>
                <button type="submit" className="submit-password-button">
                  Submit
                </button>
                {passwordChangeMessage && (
                  <div className={isSuccessMessage ? 'success-message' : 'error-message'}>
                    {passwordChangeMessage}
                  </div>
                )}
              </form>
            )}
          </div>
        </div>
      </div>
    </>
  );
}

export default MyProfile;
