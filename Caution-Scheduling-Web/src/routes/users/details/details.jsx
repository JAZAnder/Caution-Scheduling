import { useEffect } from "react";
import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import { Link, useNavigate } from 'react-router-dom';
import "../users.css"

function UserDetailsButton(user) {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const [userName, setUserName] = useState(user.user.userName)
    const [firstName, setFirstName] = useState(user.user.firstName)
    const [lastName, setLastName] = useState(user.user.lastName)
    const [email, setEmail] = useState(user.user.email)
    const [role, setRole] = useState(user.user.role)

    const updateUser = async () => {
        const myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
      
        const urlencoded = new URLSearchParams();
        urlencoded.append('userId', user.user.userId);
        urlencoded.append('userName', userName);
        urlencoded.append('firstName', firstName);
        urlencoded.append('lastName', lastName);
        urlencoded.append('email', email);
        urlencoded.append('role', role);
      
        const requestOptions = {
          method: 'PUT',
          headers: myHeaders,
          body: urlencoded,
          redirect: 'follow',
        };
      
        try {
          const response = await fetch('/api/luser/update', requestOptions);
      
          if (!response.ok) {
            throw new Error('Something Went Wrong');
          }
      
          const data = await response.json();

        } catch (error) {
            alert(error)
          console.error('Error:', error);
          throw new Error('Error');
        }
      };


    return (
        <>
            <button
                className="details-details-button"
                onClick={handleShow}
            >
                Details
            </button>

            <Modal show={show} onHide={handleClose} backdrop="static">
                <Modal.Header closeButton>
                    <Modal.Title>User Information for User #{user.user.userId}</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>UserName</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={userName}
                            onChange={(e) => setUserName(e.target.value)}
                            placeholder="UserName"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>First Name</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={firstName}
                            onChange={(e) => setFirstName(e.target.value)}
                            placeholder="First Name"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Last Name</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={lastName}
                            onChange={(e) => setLastName(e.target.value)}
                            placeholder="Last Name"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Email</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="Email"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Role</strong>
                        </Form.Label>
                        <Form.Select aria-label="Default select example" onChange={(e) => setRole(e.target.value)}>
                            <option value="">{role}</option>
                            <option value="1">Student</option>
                            <option value="2">Tutor</option>
                            <option value="3">Supervisor</option>
                            <option value="4">Administrator</option>
                            <option value="0">Disabled</option>
                        </Form.Select>
                    </Form.Group>
                </Modal.Body>
                <Modal.Footer>
                   
                        <Link to={'/user-timeslots?tutorId='+user.user.userId}> <Button variant="primary" onClick={handleClose}>Availability</Button></Link>
                    
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="danger" onClick={handleClose} disabled={true}>
                        Delete User
                    </Button>
                    <Button variant="success" className="background-1" onClick={function () { handleClose(); updateUser(); }}>
                        Save Changes
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )

}


export default UserDetailsButton