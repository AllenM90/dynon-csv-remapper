# -*- mode: python ; coding: utf-8 -*-
block_cipher = None

import os
SPECPATH = os.path.dirname(os.path.abspath(SPEC))
project_root = os.path.dirname(SPECPATH)

a = Analysis(
    [os.path.join(project_root, 'main.py')],
    pathex=[project_root, os.path.join(project_root, 'src')],
    binaries=[],
    datas=[
        (os.path.join(project_root, 'd2sCLI.exe'), '.'),
        (os.path.join(project_root, 'custom-mappings.json'), '.')
    ],
    hiddenimports=[],
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    win_no_prefer_redirects=False,
    win_private_assemblies=False,
    cipher=block_cipher,
    noarchive=False,
)
pyz = PYZ(a.pure, a.zipped_data, cipher=block_cipher)

exe = EXE(
    pyz,
    a.scripts,
    [],
    exclude_binaries=True,
    name='Dynon2Savvy',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    console=False,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
)
coll = COLLECT(
    exe,
    a.binaries,
    a.zipfiles,
    a.datas,
    strip=False,
    upx=True,
    upx_exclude=[],
    name='Dynon2Savvy',
)
