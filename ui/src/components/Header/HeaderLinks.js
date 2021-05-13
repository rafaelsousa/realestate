/*eslint-disable*/
import React from "react";
import AccountBalanceWalletOutlined from "@material-ui/icons/AccountBalanceWalletOutlined";
import IconButton from "@material-ui/core/IconButton";
// react components for routing our app without refresh
import { Link } from "react-router-dom";

// @material-ui/core components
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Tooltip from "@material-ui/core/Tooltip";

// @material-ui/icons
import { Apps, CloudDownload } from "@material-ui/icons";

// core components
import CustomDropdown from "components/CustomDropdown/CustomDropdown.js";
import Button from "components/CustomButtons/Button.js";

import styles from "assets/jss/material-kit-react/components/headerLinksStyle.js";

const useStyles = makeStyles(styles);

export default function HeaderLinks(props) {
  const classes = useStyles();
  return (
    <List className={classes.list}>
      <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText="Applications"
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={Apps}
          dropdownList={[
            <Link to="/" className={classes.dropdownLink}>
              Register a property
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
              Subscribe to rent a property
            </Link>,
          ]}
        />
      </ListItem>
        <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText="Wallet"
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={AccountBalanceWalletOutlined}
          dropdownList={[
            <Link to="/" className={classes.dropdownLink}>
              Create a new wallet
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
              Connect to an existent wallet
            </Link>,
          ]}
        />
      </ListItem>

    </List>
  );
}
