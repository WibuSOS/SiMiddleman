import Head from 'next/head';
import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import logo from './assets/logo.png'
import LoginForm from './Login';
import RegisterForm from './register';
import { useRouter } from 'next/router';

function Dashboard() {
    return (
        <div className='container mx-10 my-7'>
            Hello World
        </div>
    );
}

export default Dashboard;