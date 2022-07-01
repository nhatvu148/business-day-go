import React, { useEffect, useState } from "react";
import {
  Button,
  List,
  Typography,
  IconButton,
  styled,
  Box,
  Card,
  Tooltip,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Select,
  MenuItem,
  SelectChangeEvent,
  FormControl,
  InputLabel,
  Alert,
  AlertTitle,
} from "@mui/material";
import type { NextPage } from "next";
import ListItem from "@mui/material/ListItem";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import ListItemText from "@mui/material/ListItemText";
import Avatar from "@mui/material/Avatar";
import AddBoxIcon from "@mui/icons-material/AddBox";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";
import { pink } from "@mui/material/colors";
import { DesktopDatePicker } from "@mui/x-date-pickers/DesktopDatePicker";
import moment from "moment";
import OpenDialogDragger from "components/Draggers/OpenDialogDragger";

const Demo = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
}));

const initialCustomHolidays = [
  { date: "2022/02/02", category: "Holiday" },
  { date: "2022/03/03", category: "Business day" },
  { date: "2022/04/04", category: "Holiday" },
  { date: "2022/05/05", category: "Business day" },
  { date: "2022/06/06", category: "Holiday" },
];

function generate(element: React.ReactElement) {
  return [0, 1, 2].map((value) =>
    React.cloneElement(element, {
      key: value,
    })
  );
}

const Home: NextPage = () => {
  const [customHolidays, setCustomHolidays] = useState(initialCustomHolidays);
  const [open, setOpen] = useState(false);
  const [category, setCategory] = useState("");
  const [date, setDate] = useState<moment.Moment | null>(null);
  const [isError, setIsError] = useState(false);

  useEffect(() => {
    if (isError) {
      setTimeout(() => {
        setIsError(false);
      }, 2000);
    }
  }, [isError]);

  const handleChange = (newValue: moment.Moment | null) => {
    setDate(newValue);
  };

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleAdd = () => {
    if (date === null || category === "") {
      setIsError(true);
      return;
    }
    if (
      customHolidays.some(
        (holiday) => holiday.date === moment(date).format("YYYY/MM/DD")
      )
    ) {
      setIsError(true);
      return;
    }
    setCustomHolidays((prev) => [
      ...prev,
      { date: moment(date).format("YYYY/MM/DD"), category },
    ]);
    setDate(null);
    // setCategory("");
  };

  const handleChangeCategory = (event: SelectChangeEvent) => {
    setCategory(event.target.value);
  };

  return (
    <div>
      {isError && (
        <Alert
          sx={{
            position: "fixed",
            left: "700px",
            zIndex: 2,
            marginTop: "-15px",
          }}
          severity="error"
          onClose={() => {
            setIsError(false);
          }}
        >
          This is a error alert — check it out!
        </Alert>
      )}
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          position: "relative",
          marginTop: 40,
        }}
      >
        <Box sx={{ width: "40%" }}>
          <Card variant="outlined" style={{ padding: 10 }}>
            <div style={{ display: "flex", justifyContent: "center" }}>
              <Typography
                sx={{ mt: 1, mb: 1, mr: 2 }}
                variant="h6"
                component="div"
              >
                Custom Holidays
              </Typography>
              <Tooltip title="Add Custom Holiday">
                <IconButton
                  edge="end"
                  aria-label="add"
                  onClick={() => {
                    handleClickOpen();
                  }}
                >
                  <AddBoxIcon color="primary" />
                </IconButton>
              </Tooltip>
            </div>

            <Demo>
              <List dense={false}>
                {customHolidays.map((day, id: number) => {
                  return (
                    <ListItem
                      key={id}
                      secondaryAction={
                        <>
                          <Tooltip title="Edit">
                            <IconButton
                              edge="end"
                              aria-label="edit"
                              sx={{ mr: 0.1 }}
                            >
                              <EditIcon color="success" />
                            </IconButton>
                          </Tooltip>
                          <Tooltip title="Delete">
                            <IconButton edge="end" aria-label="delete">
                              <DeleteIcon sx={{ color: pink[500] }} />
                            </IconButton>
                          </Tooltip>
                        </>
                      }
                    >
                      <ListItemAvatar>
                        <Avatar>
                          <CalendarMonthIcon color="primary" />
                        </Avatar>
                      </ListItemAvatar>
                      <ListItemText
                        primary={day.date}
                        secondary={day.category}
                      />
                    </ListItem>
                  );
                })}
              </List>
            </Demo>
          </Card>
        </Box>
        <Dialog
          open={open}
          onClose={handleClose}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
          hideBackdrop
          disableScrollLock
          disableEnforceFocus
          PaperComponent={OpenDialogDragger}
        >
          <DialogTitle id="alert-dialog-title" sx={{ cursor: "move" }}>
            {"Add a new Holiday or Business day"}
          </DialogTitle>
          <DialogContent>
            <div
              style={{
                marginTop: 5,
                display: "flex",
                justifyContent: "center",
              }}
            >
              <DesktopDatePicker
                label="Date"
                inputFormat="YYYY-MM-DD"
                value={date}
                onChange={handleChange}
                renderInput={(params) => <TextField {...params} />}
              />
              <FormControl fullWidth sx={{ width: "50%", ml: 2 }}>
                <InputLabel id="demo-simple-select-label">Category</InputLabel>
                <Select
                  labelId="demo-simple-select-label"
                  id="demo-simple-select"
                  value={category}
                  label="Category"
                  onChange={handleChangeCategory}
                >
                  <MenuItem value={"Holiday"}>Holiday</MenuItem>
                  <MenuItem value={"Business day"}>Business day</MenuItem>
                </Select>
              </FormControl>
            </div>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleAdd}>Add</Button>
            <Button onClick={handleClose} autoFocus>
              Cancel
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    </div>
  );
};

export default Home;
