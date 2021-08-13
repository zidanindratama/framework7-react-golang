import React, { Component } from "react";
import {
  Page,
  Navbar,
  NavLeft,
  NavTitle,
  NavTitleLarge,
  NavRight,
  Link,
  Toolbar,
  Block,
  BlockTitle,
  List,
  ListItem,
  Row,
  Col,
  Button,
} from "framework7-react";
import Percobaan from "./Percobaan";

const HomePage = () => (
  <Page name="home">
    {/* Top Navbar */}
    <Navbar>
      <NavTitle sliding>Mahasiswa</NavTitle>
    </Navbar>
    {/* Toolbar */}
    {/* Page content */}
    <Block strong>
      <p>Berikut adalah daftar mahasiswa prodi informatika</p>
      <List>
        <ListItem link="/form/" title="Tambah Mahasiswa" />
      </List>
    </Block>
    <List>
      {/* <ListItem link="/about/" title="About" />
      <ListItem link="/form/" title="Form" /> */}
      <Percobaan />
    </List>
  </Page>
);
export default HomePage;
