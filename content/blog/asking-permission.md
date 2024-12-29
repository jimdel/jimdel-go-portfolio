---
title: "Permissions on iOS & Android with React Native"
blurb: "Setting up permissions on iOS & Android with react-native-permissions"
date: "2023-03-28"
---

# Mobile Permissions in React Native: A Setup Guide

Mobile permissions are essential security features that give users control over which hardware and software components an app can access. This guide explains how to implement permissions in React Native using the `react-native-permissions` library.

## Understanding Mobile Permissions

Mobile apps need explicit user permission to access device resources like:

- Camera
- Photo library
- Microphone
- Location services

Each platform handles permissions differently:

### iOS

Permissions are defined in the `Info.plist`
Example for camera access:

```xml
<key>NSCameraUsageDescription</key>
<string>This app needs camera access to take photos</string>
```

### Android

Permissions are defined in the `AndroidManifest.xml`
Example for camera access:

```xml
<uses-permission android:name="android.permission.CAMERA" />
```

### Step 1: Install react-native-permissions & pods

```bash
npm install react-native-permissions && npx pod-install
```

Note - you may need to run `react-native link react-native-permissions` if using React Native < .60 where auto-linking is not the default behavior.

### Step 2: Edit the necessary files

#### Info.plist

Add the following to your `Info.plist`:

```xml
<key>NSPhotoLibraryUsageDescription</key>
<string>$(PRODUCT_NAME) would like to access your photo library</string>

<key>NSCameraUsageDescription</key>
<string>$(PRODUCT_NAME) would like to use your camera</string>

<key>NSMicrophoneUsageDescription</key>
<string>$(PRODUCT_NAME) would like to use your microphone</string>
```

#### AndroidManifest.xml

Add the following to your `AndroidManifest.xml`:

```xml
<uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE" />
<uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE" />
<uses-permission android:name="android.permission.CAMERA" />
<uses-permission android:name="android.permission.RECORD_AUDIO" />
```

#### Podfile

Add the following to your `Podfile`:

```ruby
pod 'Permission-Camera', :path => '../node_modules/react-native-permissions/ios/Camera/Permission-Camera.podspec'
pod 'Permission-PhotoLibrary', :path => '../node_modules/react-native-permissions/ios/PhotoLibrary/Permission-PhotoLibrary.podspec'
pod 'Permission-MediaLibrary', :path => '../node_modules/react-native-permissions/ios/MediaLibrary/Permission-MediaLibrary.podspec'
pod 'Permission-PhotoLibraryAddOnly', :path => '../node_modules/react-native-permissions/ios/PhotoLibraryAddOnly/Permission-PhotoLibraryAddOnly.podspec'

```

#### package.json

Add the following to your `package.json`:

```json
"reactNativePermissionsIOS": [
    "MediaLibrary",
    "PhotoLibrary",
    "Camera",
    "PhotoLibraryAddOnly"
],
```

### Step 3: Create a usePermission hook

Create a custom React hook named `usePermission` with the following:

```ts
import { useState, useEffect } from "react";
import { Platform } from "react-native";
import {
  PERMISSIONS,
  checkMultiple,
  requestMultiple,
} from "react-native-permissions";

const usePermissions = () => {
  const [cameraStatus, setCameraStatus] = useState("");
  const [photoLibraryStatus, setPhotoLibraryStatus] = useState("");
  const [mediaLibraryStatus, setMediaLibraryStatus] = useState("");
  const [photoLibraryAddOnlyStatus, setPhotoLibraryAddOnlyStatus] =
    useState("");

  useEffect(() => {
    const requestPermissions = async () => {
      try {
        const permissions =
          Platform.OS === "android"
            ? [
                PERMISSIONS.ANDROID.CAMERA,
                PERMISSIONS.ANDROID.READ_EXTERNAL_STORAGE,
                PERMISSIONS.ANDROID.WRITE_EXTERNAL_STORAGE,
                PERMISSIONS.ANDROID.READ_EXTERNAL_STORAGE,
              ]
            : [
                PERMISSIONS.IOS.CAMERA,
                PERMISSIONS.IOS.PHOTO_LIBRARY,
                PERMISSIONS.IOS.MEDIA_LIBRARY,
                PERMISSIONS.IOS.PHOTO_LIBRARY_ADD_ONLY,
              ];
        const statuses = await checkMultiple(permissions);
        const shouldRequest = Object.values(statuses).some(
          (status) => status === "blocked" || status === "denied"
        );
        if (shouldRequest) {
          const results = await requestMultiple(permissions);
          setCameraStatus(results[permissions[0]]);
          setPhotoLibraryStatus(results[permissions[1]]);
          setMediaLibraryStatus(results[permissions[2]]);
          setPhotoLibraryAddOnlyStatus(results[permissions[3]]);
        } else {
          setCameraStatus(statuses[permissions[0]]);
          setPhotoLibraryStatus(statuses[permissions[1]]);
          setMediaLibraryStatus(statuses[permissions[2]]);
          setPhotoLibraryAddOnlyStatus(statuses[permissions[3]]);
        }
      } catch (error) {
        console.log(error);
      }
    };

    requestPermissions();
  }, []);

  return {
    cameraStatus,
    photoLibraryStatus,
    mediaLibraryStatus,
    photoLibraryAddOnlyStatus,
  };
};

export default usePermissions;
```

### Step 4 - Use in the root of your app (or where you need it)

```ts
function App(): JSX.Element {
    const {
        cameraStatus,
        photoLibraryStatus,
        mediaLibraryStatus,
        photoLibraryAddOnlyStatus,
    } = usePermissions();
    console.log({
        permissions: {
        cameraStatus,
        photoLibraryStatus,
        mediaLibraryStatus,
        photoLibraryAddOnlyStatus,
        },
    });
    return (...)
```

## Conclusion

I hope this guide helped! Please remember to do the following if you're running into build errors:

- Remove and reinstall `node_modules`
- Remove stale build crap `rm -rf path/to/ios/build` & rebuild with Xcode/Android Studio
- Clear pod cache `pod cache clean --all`
- Reinstall pods `npx pod-install`
- Check syntax of Info.plist & Android.xml
