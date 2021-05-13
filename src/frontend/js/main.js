import React from 'react'

import ReactDOM from 'react-dom';

import {
  Container,
  Row
} from 'react-bootstrap';

import Header from './components/header.js';

import {
  Tab,
  CodeArea,
  Editor
} from './components/main_editor.js';

var myCSS = require('../styles/main.css')

var bootstrapCSS = require('bootstrap/dist/css/bootstrap.css')

const App = () => {
  return (
    <Container fluid id="root_containter">
      <Row id="header_row">
        <Header />
      </Row>
      <Row id="viewport_row">
        viewport
      </Row>
      <Row id="footer_row">
        footer
      </Row>
      <div id="floater" style={{display: "none"}}></div>
    </Container>
  );
};

ReactDOM.render(App(), document.getElementById("root"));
