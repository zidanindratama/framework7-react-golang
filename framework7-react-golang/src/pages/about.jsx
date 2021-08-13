import React, { Component, useEffect, useState } from "react";
import WithListLoading from "../components/withListLoading";
import axios from "axios";
import {
  Page,
  Navbar,
  List,
  ListInput,
  ListItem,
  Toggle,
  BlockTitle,
  Row,
  Button,
  Range,
  Block,
  Col,
} from "framework7-react";

const endpoint = `http://localhost:8080/mahasiswa`;

const AboutPage = (props) => {
  const deleteTask = (nim) => {
    axios
      .delete(endpoint + "/" + nim, {
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          Origin: "http://localhost:3000/",
        },
      })
      .then((res) => {
        console.log(res.json());
      });
  };

  return (
    <>
      <Page>
        <Navbar title="Detail Mahasiswa" backLink="Back" />
        <BlockTitle>Tentang Mahasiswa</BlockTitle>
        <Block strong>
          <p>NIM : {props.nim}</p>
          <p>NAMA : {props.nama}</p>
        </Block>
        </Row>
      </Page>
    </>
  );
};

export default AboutPage;
