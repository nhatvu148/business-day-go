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
  Divider,
} from "@mui/material";
import type { NextPage } from "next";
import ListItem from "@mui/material/ListItem";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import ListItemText from "@mui/material/ListItemText";
import Avatar from "@mui/material/Avatar";
import AddBoxIcon from "@mui/icons-material/AddBox";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import CloseIcon from "@mui/icons-material/Close";
import SaveIcon from "@mui/icons-material/Save";
import CancelIcon from "@mui/icons-material/Cancel";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";
import { pink } from "@mui/material/colors";
import { DesktopDatePicker } from "@mui/x-date-pickers/DesktopDatePicker";
import moment from "moment";
import Image from "next/image";
import OpenDialogDragger from "components/Draggers/OpenDialogDragger";

const Demo = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
}));

const initialCustomHolidays = [
  { date: "2022-02-02", category: "Holiday" },
  { date: "2022-03-03", category: "Business day" },
  { date: "2022-04-04", category: "Holiday" },
  { date: "2022-05-05", category: "Business day" },
  { date: "2022-06-06", category: "Holiday" },
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
  const [updatedCategory, setUpdatedCategory] = useState("");
  const [newDate, setNewDate] = useState<moment.Moment | null>(null);
  const [newUpdatedDate, setNewUpdatedDate] = useState<moment.Moment | null>(
    null
  );
  const [isError, setIsError] = useState(false);
  const [isSuccess, setIsSuccess] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [updatedDate, setUpdatedDate] = useState("");

  useEffect(() => {
    if (isError) {
      setTimeout(() => {
        setIsError(false);
      }, 2000);
    }
  }, [isError]);

  useEffect(() => {
    if (isSuccess) {
      setTimeout(() => {
        setIsSuccess(false);
      }, 2000);
    }
  }, [isSuccess]);

  useEffect(() => {
    if (updatedDate !== "") {
      const _newUpdatedDate = moment(updatedDate);
      const _updatedCategory = customHolidays.find(
        (day) => day.date === updatedDate
      )?.category as string;
      setNewUpdatedDate(_newUpdatedDate);
      setUpdatedCategory(_updatedCategory);
    }
  }, [updatedDate, customHolidays]);

  const handleNewDateChange = (newValue: moment.Moment | null) => {
    if (newValue !== null) {
      setNewDate(newValue);
    }
  };

  const handleUpdatedDateChange = (newValue: moment.Moment | null) => {
    if (newValue !== null) {
      setNewUpdatedDate(newValue);
    }
  };

  const handleClickOpen = () => {
    setOpen(true);
  };

  const closeDialog = () => {
    setOpen(false);
  };

  const handleClose = (
    event: {},
    reason: "backdropClick" | "escapeKeyDown"
  ) => {
    if (reason && reason == "backdropClick") {
      return;
    }
    closeDialog();
  };

  const handleAdd = () => {
    if (newDate === null || category === "") {
      setErrorMessage("Date or Category cannot be empty");
      setIsError(true);
      return;
    }
    const _newDate = moment(newDate).format("YYYY-MM-DD");
    if (customHolidays.some((holiday) => holiday.date === _newDate)) {
      setErrorMessage(`Date ${_newDate} already exists`);
      setIsError(true);
      return;
    }
    setCustomHolidays((prev) => [
      ...prev,
      { date: moment(newDate).format("YYYY-MM-DD"), category },
    ]);
    setNewDate(null);
    // setCategory("");
  };

  const handleChangeCategory = (event: SelectChangeEvent) => {
    setCategory(event.target.value);
  };

  const handleChangeUpdatedCategory = (event: SelectChangeEvent) => {
    setUpdatedCategory(event.target.value);
  };

  const updateHoliday = () => {
    if (newUpdatedDate === null || updatedCategory === "") {
      setErrorMessage("Date or Category cannot be empty");
      setIsError(true);
      return;
    }
    const _newDate = moment(newUpdatedDate).format("YYYY-MM-DD");
    if (
      customHolidays
        .filter((holiday) => holiday.date !== updatedDate)
        .some((holiday) => holiday.date === _newDate)
    ) {
      setErrorMessage(`Date ${_newDate} already exists`);
      setIsError(true);
      return;
    }
    setCustomHolidays((prev) =>
      prev.map((holiday) => {
        if (holiday.date === updatedDate) {
          return {
            date: newUpdatedDate.format("YYYY-MM-DD"),
            category: updatedCategory,
          };
        } else {
          return holiday;
        }
      })
    );

    cancelEdit();
    setSuccessMessage("Saved");
    setIsSuccess(true);
  };

  const cancelEdit = () => {
    setNewUpdatedDate(null);
    setUpdatedCategory("");
    setUpdatedDate("");
  };

  return (
    <div>
      {isError && (
        <Alert
          sx={{
            position: "fixed",
            left: "700px",
            zIndex: 2,
            marginTop: "-40px",
          }}
          severity="error"
          onClose={() => {
            setIsError(false);
          }}
        >
          {errorMessage}
        </Alert>
      )}
      {isSuccess && (
        <Alert
          sx={{
            position: "fixed",
            left: "700px",
            zIndex: 2,
            marginTop: "-40px",
          }}
          onClose={() => {
            setIsSuccess(false);
          }}
        >
          {successMessage}
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
                {customHolidays.length === 0 && (
                  <div
                    style={{
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <div style={{ position: "absolute", top: 20 }}>No data</div>
                    <Image
                      src="/empty.svg"
                      alt="No data"
                      width={100}
                      height={80}
                    />
                    <Divider />
                  </div>
                )}

                {customHolidays
                  .sort(
                    (a, b) =>
                      moment(a.date).valueOf() - moment(b.date).valueOf()
                  )
                  .map((day, id: number) => {
                    if (day.date === updatedDate) {
                      return (
                        <div>
                          <div
                            style={{
                              marginTop: 5,
                              display: "flex",
                              justifyContent: "center",
                            }}
                          >
                            <Avatar sx={{ mr: 3, mt: 2, mb: 2 }}>
                              <CalendarMonthIcon color="primary" />
                            </Avatar>
                            <DesktopDatePicker
                              label="Date"
                              inputFormat="YYYY-MM-DD"
                              value={newUpdatedDate}
                              onChange={handleUpdatedDateChange}
                              renderInput={(params) => (
                                <TextField {...params} />
                              )}
                            />
                            <FormControl
                              fullWidth
                              sx={{ width: "34%", ml: 2, mr: 2 }}
                            >
                              <InputLabel id="demo-simple-select-label">
                                Category
                              </InputLabel>
                              <Select
                                labelId="demo-simple-select-label"
                                id="demo-simple-select"
                                value={updatedCategory}
                                label="Category"
                                onChange={handleChangeUpdatedCategory}
                              >
                                <MenuItem value={"Holiday"}>Holiday</MenuItem>
                                <MenuItem value={"Business day"}>
                                  Business day
                                </MenuItem>
                              </Select>
                            </FormControl>
                            <Tooltip title={"Save"}>
                              <IconButton
                                aria-label="save"
                                onClick={() => {
                                  updateHoliday();
                                }}
                              >
                                <SaveIcon color="primary" />
                              </IconButton>
                            </Tooltip>
                            <Tooltip title={"Cancel"}>
                              <IconButton
                                aria-label="cancel"
                                onClick={() => {
                                  cancelEdit();
                                }}
                              >
                                <CancelIcon sx={{ color: pink[500] }} />
                              </IconButton>
                            </Tooltip>
                          </div>
                          <Divider />
                        </div>
                      );
                    }
                    return (
                      <div key={day.date}>
                        <ListItem
                          secondaryAction={
                            <>
                              <Tooltip title="Edit">
                                <IconButton
                                  edge="end"
                                  aria-label="edit"
                                  sx={{ mr: 0.1 }}
                                  onClick={() => {
                                    setUpdatedDate(day.date);
                                  }}
                                >
                                  <EditIcon color="success" />
                                </IconButton>
                              </Tooltip>
                              <Tooltip title="Delete">
                                <IconButton
                                  edge="end"
                                  aria-label="delete"
                                  onClick={() => {
                                    setCustomHolidays((prev) =>
                                      prev.filter(
                                        (holiday) => holiday.date !== day.date
                                      )
                                    );
                                  }}
                                >
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
                        <Divider />
                      </div>
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
          // hideBackdrop
          disableScrollLock
          disableEnforceFocus
          // @ts-ignore
          PaperComponent={OpenDialogDragger}
          sx={{ pointerEvents: "auto" }}
        >
          <DialogTitle id="alert-dialog-title" sx={{ cursor: "move" }}>
            <div style={{ marginRight: 25 }}>
              {"Add a new Holiday or Business day"}
            </div>
            <IconButton
              aria-label="close"
              onClick={closeDialog}
              sx={{
                position: "absolute",
                right: 8,
                top: 12,
                color: (theme) => theme.palette.grey[500],
              }}
            >
              <CloseIcon />
            </IconButton>
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
                value={newDate}
                onChange={handleNewDateChange}
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
            <Button onClick={closeDialog}>Cancel</Button>
          </DialogActions>
        </Dialog>
      </div>
    </div>
  );
};

export default Home;
